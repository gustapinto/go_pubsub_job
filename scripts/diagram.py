import diagrams as d
import diagrams.programming.language as pl
import diagrams.gcp.analytics as ga


with d.Diagram(name="Services", filename="services", curvestyle="curved"):
    publisher = pl.Go(label="Publisher", nodeid="publisher")

    with d.Cluster(label="Google Cloud"):
        pubsub = ga.PubSub(label="Pub/Sub Jobs")
        bigquery = ga.BigQuery(label="Big Query Results")

    consumer = pl.Go(label="Consumer", nodeid="consumer")

    publisher >> pubsub >> consumer << bigquery
