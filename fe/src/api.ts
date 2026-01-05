import {
  AnonymousAuthenticationProvider,
  DefaultApiError,
} from "@microsoft/kiota-abstractions";
import { FetchRequestAdapter } from "@microsoft/kiota-http-fetchlibrary";
import { createMinesClient, type MinesClient } from "./client/minesClient";
import type { MoveDto, NewGameDto } from "./client/models/req";
import type { MatchstateDto } from "./client/models/res";
import { failure, type Result, success } from "./result";

export interface MinesApiClient {
  fetchMatch(matchId: string): Promise<Result<MatchstateDto>>;
  startNewGame(newGame: NewGameDto): Promise<Result<MatchstateDto>>;
  makeMove(matchId: string, move: MoveDto): Promise<Result<MatchstateDto>>;
}

export function makeNewApiClient(baseUrl: string): MinesApiClient {
  const authProvider = new AnonymousAuthenticationProvider();

  const fetchAdapter = new FetchRequestAdapter(authProvider);

  fetchAdapter.baseUrl = baseUrl;

  const kiotaClient = createMinesClient(fetchAdapter);

  return {
    fetchMatch: (matchId: string) => fetchMatch(kiotaClient, matchId),
    startNewGame: (newGameRequest: NewGameDto) =>
      startNewGame(kiotaClient, newGameRequest),
    makeMove: (matchId: string, move: MoveDto) =>
      makeMove(kiotaClient, matchId, move),
  };
}

async function fetchMatch(
  kiotaClient: MinesClient,
  matchId: string,
): Promise<Result<MatchstateDto>> {
  return safeApiCall(
    () => kiotaClient.match.byMatchId(matchId).get(),
    `Failed to fetch match ${matchId}`,
  );
}

async function startNewGame(
  kiotaClient: MinesClient,
  newGameRequest: NewGameDto,
): Promise<Result<MatchstateDto>> {
  return safeApiCall(
    () => kiotaClient.match.post(newGameRequest),
    "Failed to start new game",
  );
}

async function makeMove(
  kiotaClient: MinesClient,
  matchId: string,
  move: MoveDto,
): Promise<Result<MatchstateDto>> {
  return safeApiCall(
    () => kiotaClient.match.byMatchId(matchId).move.post(move),
    `Failed to make a move on ${matchId}`,
  );
}

async function safeApiCall<T>(
  apiCall: () => Promise<T | undefined>,
  errorMessage: string,
): Promise<Result<T>> {
  try {
    const response = await apiCall();

    if (response === undefined) {
      return failure({ message: errorMessage });
    }

    return success(response);
  } catch (maybeApiError: unknown) {
    if (maybeApiError instanceof DefaultApiError) {
      return failure({
        message: errorMessage,
        cause: maybeApiError,
      });
    } else {
      return failure({
        message: `${errorMessage} in general`,
        cause: maybeApiError as Error,
      });
    }
  }
}
