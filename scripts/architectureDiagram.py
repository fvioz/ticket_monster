#!/usr/bin/env python3

from diagrams import Cluster, Diagram, Edge
from diagrams.onprem.compute import Server
from diagrams.onprem.database import PostgreSQL
from diagrams.onprem.inmemory import Redis
from diagrams.onprem.network import Nginx
from diagrams.programming.language import Go

graph_attr = {"bgcolor": "white"}

with Diagram(
    "Solution in scale",
    filename="doc/images/architecture",
    show=False,
    graph_attr=graph_attr,
    direction="TB",
):
    ingress = Nginx("ingress")
    event_fetcher = Go("Event Fetcher")

    with Cluster("API Cluster"):
        api = [
            Server("API 1"),
            Server("API 2"),
            Server("API 3"),
        ]

    with Cluster("Database HA"):
        db_primary = PostgreSQL("Primary")
        (
            db_primary
            - Edge(color="darkblue", style="dotted")
            - [
                PostgreSQL("Replica 1"),
                PostgreSQL("Replica 2"),
                PostgreSQL("Replica N"),
            ]
        )

    with Cluster("Procesors Cluster"):
        event_processor = [
            Go("Event processor 1"),
            Go("Event processor 2"),
            Go("Event processor 3"),
        ]

    with Cluster("Redis HA"):
        redis_primary = Redis("Primary")
        (
            redis_primary
            - Edge(color="red", style="dotted")
            - [
                Redis("Replica 1"),
                Redis("Replica 2"),
                Redis("Replica N"),
            ]
        )

    event_fetcher >> Edge(color="purple") << redis_primary
    event_processor >> Edge(color="orange") >> db_primary
    event_processor >> Edge(color="black") << redis_primary

    api << Edge(color="orange") << db_primary
    ingress >> Edge(color="darkgreen") << api
