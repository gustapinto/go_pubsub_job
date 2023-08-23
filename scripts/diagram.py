import diagrams as d
import diagrams.programming.language as pl
import diagrams.gcp.analytics as ga
import diagrams.gcp.database as gd


with d.Diagram(name="Services", filename="services", curvestyle="curved"):
    publisher = pl.Go(label="Publisher", nodeid="publisher")

    with d.Cluster(label="Google Cloud Platform"):
        pubsub = ga.PubSub(label="Pub/Sub Jobs")
        firestore = gd.Firestore(label="Firestore Job States")

    consumer = pl.Go(label="Consumer", nodeid="consumer")

    publisher \
        >> d.Edge(forward=True, label='Envia Jobs Para Mensageria') \
        >> pubsub \
        >> d.Edge(reverse=True, label='Consome Jobs da Mensageria') \
        >> consumer \
        >> d.Edge(forward=True, reverse=True, label='Pesiste Estado dos Jobs') \
        >> firestore
