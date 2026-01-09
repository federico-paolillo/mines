import { type ComponentChildren, createContext } from "preact";
import { useContext, useRef } from "preact/hooks";
import { type MinesApiClient, makeNewApiClient } from "./api";

const ClientContext = createContext<MinesApiClient | undefined>(undefined);

interface ClientProviderProps {
  children: ComponentChildren;
  apiBaseUrl: string;
}

export function ClientProvider({ children, apiBaseUrl }: ClientProviderProps) {
  const clientRef = useRef<MinesApiClient>(makeNewApiClient(apiBaseUrl));

  return (
    <ClientContext.Provider value={clientRef.current}>
      {children}
    </ClientContext.Provider>
  );
}

export function useApiClient(): MinesApiClient {
  const context = useContext(ClientContext);

  if (context === undefined) {
    throw new Error("useApiClient must be used within a ClientProvider");
  }

  return context;
}
