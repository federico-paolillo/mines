import { useEffect, useState } from "preact/hooks";
import { useLocation } from "preact-iso";
import { MovetypeObject } from "../client/models/matchmaking";
import type { MatchstateDto } from "../client/models/res";
import { useApiClient } from "../clientContext";

export function useGameState(gameId: string | null) {
  const { route } = useLocation();
  const client = useApiClient();

  const [gameState, setGameState] = useState<MatchstateDto | null>(null);
  const [loading, setLoading] = useState(true);

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
      }
    };

    fetchGame();
  }, [gameId, client, route]);

  const handleMove = async (x: number, y: number, type: MovetypeObject) => {
    if (!gameId) return;

    setLoading(true);

    const result = await client.makeMove(gameId, {
      x,
      y,
      type,
    });

    setLoading(false);

    if (result.success) {
      setGameState(result.value);
      if (result.value.lives === 0) {
        route("/game-over");
      }
    } else if (result.error.kind === "match_over") {
      route("/game-over");
    }
  };

  const onOpenCell = async (x: number, y: number) => {
    await handleMove(x, y, MovetypeObject.Open);
  };

  const onFlagCell = async (x: number, y: number) => {
    await handleMove(x, y, MovetypeObject.Flag);
  };

  const onExpired = () => {
    route("/game-over");
  };

  return {
    gameState,
    loading,
    onOpenCell,
    onFlagCell,
    onExpired,
  };
}
