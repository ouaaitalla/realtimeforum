import { apiFetch } from "../utils/fetch.js";

export function getCategoriesRequest() {
    return apiFetch("/categories", {
        method: "GET",
    });
}