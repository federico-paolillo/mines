import { DefaultApiError } from "@microsoft/kiota-abstractions";
import { useEffect, useState } from "preact/hooks";
import { useLocation } from "preact-iso";
import { MovetypeObject } from "../client/models/matchmaking";
import type { MatchstateDto } from "../client/models/res";
import { useApiClient } from "../clientContext";

export function useGameState(gameId: string | null) {
  const { route } = useLocation();
  const client = useApiClient();

  const [gameState, setGameState] = useState<MatchstateDto | null>(null);
  const [loading, setLoading] = useState(!!gameId);

  useEffect(() => {
    if (!gameId) {
      setLoading(false);
      return;
    }

    const fetchGame = async () => {
      setLoading(true);

      const result = await client.fetchMatch(gameId);

      setLoading(false);

      if (result.success) {
        if (result.value.lives === 0) {
          route("/game-over");
          return;
        }
        setGameState(result.value);
      } else {
        console.error(result.error);
      }
    };

    fetchGame();
  }, [gameId, client, route]);

  const handleCellClick = async (x: number, y: number) => {
    if (!gameId) return;

    const result = await client.makeMove(gameId, {
      x,
      y,
      type: MovetypeObject.Open,
    });

    if (result.success) {
      setGameState(result.value);
      if (result.value.lives === 0) {
        route("/game-over");
      }
    } else {
      if (
        result.error.cause instanceof DefaultApiError &&
        result.error.cause.responseStatusCode === 422
      ) {
        route("/game-over");
      } else {
        console.error(result.error);
      }
    }
  };

  const handleCellRightClick = async (x: number, y: number) => {
    if (!gameId) return;

    const result = await client.makeMove(gameId, {
      x,
      y,
      type: MovetypeObject.Flag,
    });

    if (result.success) {
      setGameState(result.value);
      if (result.value.lives === 0) {
        route("/game-over");
      }
    } else {
      if (
        result.error.cause instanceof DefaultApiError &&
        result.error.cause.responseStatusCode === 422
      ) {
        route("/game-over");
      } else {
        console.error(result.error);
      }
    }
  };

  const handleExpired = () => {
    route("/game-over");
  };

  return {
    gameState,
    loading,
    onCellClick: handleCellClick,
    onCellRightClick: handleCellRightClick,
    onExpired: handleExpired,
  };
}
