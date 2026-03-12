import { DefaultApiError } from "@microsoft/kiota-abstractions";
import { type ComponentChildren, createContext } from "preact";
import { useContext, useRef } from "preact/hooks";
import { type MinesApiClient, makeNewApiClient } from "./api";
import type { Result } from "./result";

export const ClientContext = createContext<MinesApiClient | undefined>(
  undefined,
);

interface ClientProviderProps {
  children: ComponentChildren;
  apiBaseUrl: string;
}

export function ClientProvider({ children, apiBaseUrl }: ClientProviderProps) {
  const clientRef = useRef<MinesApiClient>(
    wrapClient(makeNewApiClient(apiBaseUrl)),
  );

  return (
    <ClientContext.Provider value={clientRef.current}>
      {children}
    </ClientContext.Provider>
  );
}

export function wrapClient(client: MinesApiClient): MinesApiClient {
  async function wrapCall<T>(call: Promise<Result<T>>): Promise<Result<T>> {
    const result = await call;
    if (!result.success) {
      if (
        result.error.cause instanceof DefaultApiError &&
        result.error.cause.responseStatusCode === 422
      ) {
        result.error.kind = "match_over";
      }
      console.error(result.error.message, result.error.cause);
    }
    return result;
  }

  return {
    fetchMatch: (matchId) => wrapCall(client.fetchMatch(matchId)),
    startNewGame: (newGame) => wrapCall(client.startNewGame(newGame)),
    makeMove: (matchId, move) => wrapCall(client.makeMove(matchId, move)),
  };
}

export function useApiClient(): MinesApiClient {
  const context = useContext(ClientContext);

  if (context === undefined) {
    throw new Error("useApiClient must be used within a ClientProvider");
  }

  return context;
}
