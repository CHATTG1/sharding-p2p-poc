docker build --build-arg GIT_COMMIT=$(git rev-list -1 HEAD) -t ethresearch/sharding-p2p:dev .