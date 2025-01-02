export type ApiError = {
	errors?: string[];
};

export function isApiError(obj: any): obj is ApiError {
	return obj != null && (obj as ApiError).errors !== undefined;
}
