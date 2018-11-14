from datetime import (
    datetime,
)
import json
import re
import subprocess
import sys
import threading
import os
import time

from dateutil import (
    parser,
)


RPC_PORT_BASE = 13000
PORT_BASE = 10000


def get_docker_host_ip():
    sysname = os.uname().sysname
    if sysname != 'Darwin' and sysname != 'Linux':
        raise ValueError(
            "Failed to get ip in platforms other than Linux and macOS: {}".format(sysname)
        )
    cmd = 'ifconfig | grep -E "([0-9]{1,3}\\.){3}[0-9]{1,3}" | grep -v 127.0.0.1 | awk \'{ print $2 }\' | cut -f2 -d: | head -n1'
    res = subprocess.run(cmd, shell=True, check=True, stdout=subprocess.PIPE, encoding='utf-8')
    return res.stdout.rstrip()


class CLIFailure(Exception):
    pass


class Node:

    ip = None
    port = None
    rpc_port = None
    peer_id = None
    seed = None

    def __init__(self, ip, port, rpc_port, seed):
        self.ip = ip
        self.port = port
        self.rpc_port = rpc_port
        self.seed = seed

    def __repr__(self):
        return "<Node seed={} peer_id={}>".format(
            self.seed,
            None if self.peer_id is None else self.peer_id[2:8],
        )

    @property
    def name(self):
        return f"whiteblock-node{self.seed}"

    @property
    def multiaddr(self):
        return f"/ip4/{self.ip}/tcp/{self.port}/ipfs/{self.peer_id}"

    def close(self):
        subprocess.run(f"docker kill {self.name}", shell=True, stdout=subprocess.PIPE)
        subprocess.run(f"docker rm -f {self.name}", shell=True, stdout=subprocess.PIPE)

    def run(self, bootnodes=None):
        """`bootnodes` should be a list of string. Each string should be a multiaddr.
        """
        self.close()
        bootnodes_cmd = ""
        if bootnodes is not None:
            bootnodes_cmd = "-bootstrap -bootnodes={}".format(
            ",".join(bootnodes),
        )
        cmd = "docker run -d --name {} -p {}:10000 -p {}:13000 ethresearch/sharding-p2p:dev sh -c \"./sharding-p2p-poc -loglevel=DEBUG -ip=0.0.0.0 -seed={} {}\"".format(
            self.name,
            self.port,
            self.rpc_port,
            self.seed,
            bootnodes_cmd,
        )
        subprocess.run(cmd, shell=True, stdout=subprocess.PIPE, check=True)

    def cli(self, cmd, **kwargs):
        cmd_list = cmd.split(' ')
        cmd_quoted_param_list = ["'{}'".format(i) for i in cmd_list]
        cmd_quoted_param_str = " ".join(cmd_quoted_param_list)
        return subprocess.run(
            [
                "docker",
                "exec",
                "-t",
                self.name,
                "sh",
                "-c",
                "./sharding-p2p-poc '-client' {}".format(cmd_quoted_param_str),
            ],
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            encoding='utf-8',
            **kwargs,
        )

    def cli_safe(self, cmd):
        res = self.cli(cmd)
        if res.returncode != 0:
            print(res)
            raise CLIFailure(
                "exit_code={}, stdout={!r}, stderr={!r}".format(
                    res.returncode,
                    res.stdout,
                    res.stderr,
                )
            )
        out = res.stdout.rstrip()
        # assume CLIs only reply data in JSON
        if out != '':
            return json.loads(out)
        return out

    def add_peer(self, node):
        self.cli_safe("addpeer {} {} {}".format(node.ip, node.port, node.seed))

    def remove_peer(self, peer_id):
        self.cli_safe("removepeer {}".format(peer_id))

    def list_peer(self):
        return self.cli_safe("listpeer")

    def list_topic_peer(self, topics):
        return self.cli_safe("listtopic {}".format(' '.join(topics)))

    def subscribe_shard(self, shard_ids):
        return self.cli_safe("subshard {}".format(' '.join(map(str, shard_ids))))

    def unsubscribe_shard(self, shard_ids):
        return self.cli_safe("unsubshard {}".format(' '.join(map(str, shard_ids))))

    def get_subscribed_shard(self):
        return self.cli_safe("getsubshard")

    def broadcast_collation(self, shard_id, num_collations, collation_size, collation_time):
        return self.cli_safe("broadcastcollation {} {} {} {}".format(
            shard_id,
            num_collations,
            collation_size,
            collation_time,
        ))

    def stop(self):
        return self.cli_safe("stop")

    def grep_log(self, pattern):
        res = subprocess.run(
            [
                "docker logs {} 2>&1 | grep '{}'".format(
                    self.name,
                    pattern,
                ),
            ],
            shell=True,
            stdout=subprocess.PIPE,
            encoding='utf-8',
        )
        return res.stdout.rstrip()

    def set_peer_id(self):
        grep_res = self.grep_log('Node is listening')
        match = re.search(r'peerID=([a-zA-Z0-9]+) ', grep_res)
        if match is None:
            raise ValueError("failed to grep the peer_id from docker logs")
        self.peer_id = match[1]

    def wait_for_log(self, pattern, k_th):
        """Wait for the `k_th` log in the format of `pattern`
        """
        # TODO: should set a `timeout`?
        cmd = "docker logs {} -t -f 2>&1 | grep --line-buffered '{}' -m {}".format(
            self.name,
            pattern,
            k_th + 1,
        )
        res = subprocess.run(
            [cmd],
            shell=True,
            stdout=subprocess.PIPE,
            encoding='utf-8',
        )
        logs = res.stdout.rstrip()
        log = logs.split('\n')[k_th]
        time_str_iso8601 = log.split(' ')[0]
        # # get rid of the control elements
        # match = re.search(r'\x1b\[[0-9;]+[A-Za-z]([:\.0-9]+)', time_str)
        return parser.parse(time_str_iso8601)

def make_node(seed, bootnodes=None):
    n = Node(
        get_docker_host_ip(),
        seed + PORT_BASE,
        seed + RPC_PORT_BASE,
        seed,
    )
    n.run(bootnodes)
    return n


def make_local_nodes(low, top, bootnodes=None):
    nodes = []
    threads = []

    def run_node(seed, bootnodes=None):
        node = make_node(seed, bootnodes)
        nodes.append(node)

    for i in range(low, top):
        t = threading.Thread(target=run_node, args=(i, bootnodes))
        t.start()
        threads.append(t)
    for t in threads:
        t.join()
    nodes = sorted(nodes, key=lambda node: node.seed)

    sleep_time = 5 + (top - low)
    time.sleep(sleep_time)

    threads = []
    for node in nodes:
        t = threading.Thread(target=node.set_peer_id, args=())
        t.start()
        threads.append(t)
    for t in threads:
        t.join()
    return nodes


def connect_barbell(nodes):
    threads = []
    for i in range(len(nodes) - 1):
        t = threading.Thread(target=nodes[i].add_peer, args=(nodes[i + 1],))
        t.start()
        threads.append(t)
    for t in threads:
        t.join()


def connect_fully(nodes):
    threads = []
    for i in range(len(nodes) - 1):
        for j in range(i + 1, len(nodes)):
            t = threading.Thread(target=nodes[i].add_peer, args=(nodes[j],))
            t.start()
            threads.append(t)
    for t in threads:
        t.join()


def ensure_barbell_connections(nodes):
    threads = []

    def check_connection(nodes, i):
        if len(nodes) <= 1:
            return
        peers = nodes[i].list_peer()
        if i == 0:
            assert len(peers) == 1 and peers[0] == nodes[i + 1].peer_id
        elif i == len(nodes) - 1:
            assert len(peers) == 1 and peers[0] == nodes[i - 1].peer_id
        else:
            assert len(peers) == 2 and \
                (nodes[i - 1].peer_id in peers) and (nodes[i + 1].peer_id in peers)

    for idx, _ in enumerate(nodes):
        t = threading.Thread(target=check_connection, args=(nodes, idx))
        t.start()
        threads.append(t)
    for t in threads:
        t.join()


def kill_nodes(nodes):
    node_names = [n.name for n in nodes]
    subprocess.run(
        ["docker", "kill"] + node_names,
        stdout=subprocess.PIPE,
    )
