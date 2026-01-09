import { createContext } from "preact";
import { useContext, useRef } from "preact/hooks";
import { type MinesApiClient, makeNewApiClient } from "./api";

const ClientContext = createContext<MinesApiClient | undefined>(undefined);

interface ClientProviderProps {
  children: any;
}

export function ClientProvider({ children }: ClientProviderProps) {
  // Use a dummy URL as per previous instructions since we don't have a real backend
  const clientRef = useRef<MinesApiClient>(
    makeNewApiClient("http://localhost:8080"),
  );

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
