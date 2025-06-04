import { useCallback, useState } from "react";

import { WebSocketMessage } from "@types/ws/ws";

interface WebSocketProps {
  wsPath: string;
  interval: number; // Default ping interval in milliseconds
  enabled?: boolean; // Optional, default is true
}

interface WebSocketReturn {
  position: number;
}

export const useWS = ({
  wsPath,
  interval = 30000,
  enabled = true,
}: WebSocketProps): WebSocketReturn => {
  const [position, setPossition] = useState<number>(0);
  const [webSocket, setWebSocket] = useState<WebSocket | undefined>(undefined);

  const delay = (ms: number) => new Promise((res) => setTimeout(res, ms));

  const ping = useCallback(() => {
    if (webSocket) {
      webSocket.send("PING");
    }
  }, [webSocket]);

  const connect = useCallback((): void => {
    if (webSocket) {
      return;
    }

    const ws = new WebSocket("http://localhost:8081" + wsPath);
    setWebSocket(ws);

    ws.onopen = function (evt) {
      console.info("ws:open", evt);
    };

    ws.onclose = function (evt) {
      console.info("ws:close", evt);
      setWebSocket(undefined);
    };

    ws.onmessage = function (evt) {
      if (evt.data) {
        const data =
          typeof evt.data === "string"
            ? (JSON.parse(evt.data) as WebSocketMessage)
            : { pos: 0 };

        setPossition(data.pos);
      }
    };

    ws.onerror = async function (evt) {
      console.error("ws:error", evt);
      await delay(5000);
      connect();
    };
  }, [webSocket, wsPath]);

  if (enabled && !webSocket) {
    setInterval(
      function () {
        if (webSocket && enabled) {
          ping();
        } else if (enabled) {
          connect();
        }
      },
      interval + Math.floor(Math.random() * 10000),
    );
  }

  return {
    position,
  };
};
