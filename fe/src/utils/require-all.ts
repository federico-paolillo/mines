export type RequireAll<T> = { [K in keyof T]-?: Required<NonNullable<T[K]>> };
