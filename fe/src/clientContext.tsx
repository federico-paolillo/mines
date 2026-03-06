import { DefaultApiError } from "@microsoft/kiota-abstractions";
import { type ComponentChildren, createContext } from "preact";
import { useContext, useRef } from "preact/hooks";
import { type MinesApiClient, makeNewApiClient } from "./api";
import type { MoveDto, NewGameDto } from "./client/models/req";
import type { MatchstateDto } from "./client/models/res";
import type { ApiError, Result } from "./result";

export interface GameClient {
  fetchMatch(matchId: string): Promise<Result<MatchstateDto, ApiError>>;
  startNewGame(newGame: NewGameDto): Promise<Result<MatchstateDto, ApiError>>;
  makeMove(
    matchId: string,
    move: MoveDto,
  ): Promise<Result<MatchstateDto, ApiError>>;
}

export const ClientContext = createContext<GameClient | undefined>(undefined);

interface ClientProviderProps {
  children: ComponentChildren;
  apiBaseUrl: string;
}

export function ClientProvider({ children, apiBaseUrl }: ClientProviderProps) {
  const clientRef = useRef<GameClient>(
    wrapClient(makeNewApiClient(apiBaseUrl)),
  );

  return (
    <ClientContext.Provider value={clientRef.current}>
      {children}
    </ClientContext.Provider>
  );
}

export function wrapClient(client: MinesApiClient): GameClient {
  async function wrapCall<T>(
    call: Promise<Result<T>>,
  ): Promise<Result<T, ApiError>> {
    const result = await call;
    if (result.success) {
      return result;
    }

    const apiError = toApiError(result.error);
    console.error(result.error.message, result.error.cause);
    return { success: false, error: apiError };
  }

  return {
    fetchMatch: (matchId) => wrapCall(client.fetchMatch(matchId)),
    startNewGame: (newGame) => wrapCall(client.startNewGame(newGame)),
    makeMove: (matchId, move) => wrapCall(client.makeMove(matchId, move)),
  };
}

function toApiError(problem: { message: string; cause?: Error }): ApiError {
  if (
    problem.cause instanceof DefaultApiError &&
    problem.cause.responseStatusCode === 422
  ) {
    return { kind: "match_over", message: problem.message };
  }
  return { kind: "unknown", message: problem.message };
}

export function useApiClient(): GameClient {
  const context = useContext(ClientContext);

  if (context === undefined) {
    throw new Error("useApiClient must be used within a ClientProvider");
  }

  return context;
}
