#!/usr/bin/env python3

from diagrams import Cluster, Diagram, Edge
from diagrams.onprem.compute import Server
from diagrams.onprem.client import Users
from diagrams.onprem.database import PostgreSQL
from diagrams.onprem.inmemory import Redis
from diagrams.onprem.network import Nginx
from diagrams.programming.language import Go, NodeJS

graph_attr = {"bgcolor": "white"}

with Diagram(
    "Solution in scale",
    filename="doc/images/architecture",
    show=False,
    graph_attr=graph_attr,
    direction="TB",
):
    users = Users("Users")

    with Cluster("App"):
        ingress = Nginx("ingress")
        event_fetcher = Go("Event Fetcher")

        with Cluster("Web Service Cluster"):
            web = NodeJS("ws")
            web - Edge(color="gray", style="dashed") - NodeJS("ws N")

        with Cluster("API Service Cluster"):
            api = Go("api")
            api - Edge(color="gray", style="dashed") - Go("api N")

        with Cluster("Web Socket Service Cluster"):
            web_socket = Go("ws")
            web_socket - Edge(color="gray", style="dashed") - Go("ws N")

        with Cluster("Procesors Cluster"):
            event_processor = Go("ep1")
            event_processor - Edge(color="gray", style="dashed") - Go("ep N")

        with Cluster("Database HA"):
            db_primary = PostgreSQL("storage")
            db_primary - Edge(color="darkblue", style="dashed") - PostgreSQL("replica")

            event_processor >> Edge(color="orange") >> db_primary
            api << Edge(color="orange") << db_primary
            web_socket >> Edge(color="orange") << db_primary

        with Cluster("Tasks HA"):
            redis_tasks_primary = Redis("tasks")
            (
                redis_tasks_primary
                - Edge(color="darkred", style="dashed")
                - Redis("replica")
            )

            event_fetcher >> Edge(color="purple") << redis_tasks_primary
            event_processor >> Edge(color="black") << redis_tasks_primary

        with Cluster("Web Socket HA"):
            redis_ws_primary = Redis("web socket")
            (
                redis_ws_primary
                - Edge(color="darkred", style="dashed")
                - Redis("replica")
            )

            web_socket >> Edge(color="purple") << redis_ws_primary

        ingress >> Edge(color="darkgreen") << web
        ingress >> Edge(color="darkgreen") << api
        ingress >> Edge(color="darkgreen") << web_socket

        ingress >> Edge(color="black") << users
