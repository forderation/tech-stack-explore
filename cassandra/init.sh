# Run the first node and keep it in background up and running
docker run --name cassandra-1 --network host -p 9042:9042 -d cassandra:3.11
INSTANCE1=$(docker inspect --format="{{ .NetworkSettings.IPAddress }}" cassandra-1)
echo "Instance 1: ${INSTANCE1}"

# Run the second node
docker run --name cassandra-2 --network host -p 9043:9042 -d -e CASSANDRA_SEEDS=host cassandra:3.11
INSTANCE2=$(docker inspect --format="{{ .NetworkSettings.IPAddress }}" cassandra-2)
echo "Instance 2: ${INSTANCE2}"

echo "Wait 60s until the second node joins the cluster"
sleep 60

# Run the third node
docker run --name cassandra-3 --network host -p 9044:9042 -d -e CASSANDRA_SEEDS=host cassandra:3.11
INSTANCE3=$(docker inspect --format="{{ .NetworkSettings.IPAddress }}" cassandra-3)