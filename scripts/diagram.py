import diagrams as d
import diagrams.programming.language as pl
import diagrams.gcp.analytics as ga
import diagrams.gcp.database as gd


with d.Diagram(name="Services", filename="services", curvestyle="curved"):
    publisher = pl.Go(label="Publisher", nodeid="publisher")

    with d.Cluster(label="Google Cloud Platform"):
        pubsub = ga.PubSub(label="Pub/Sub jobs queue")
        firestore = gd.Firestore(label="Firestore job status store")

    consumer = pl.Go(label="Consumer", nodeid="consumer")

    publisher \
        >> d.Edge(forward=True, label='Send new jobs to Pub/Sub') \
        >> pubsub \
        >> d.Edge(reverse=True, label='Execute job from Pub/Sub message') \
        >> consumer \
        >> d.Edge(forward=False, reverse=True, label='Save job status into Firestore') \
        >> firestore
