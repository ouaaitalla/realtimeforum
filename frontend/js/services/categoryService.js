import { getCategoriesRequest } from "../api/categories.js";

export async function getCategories() {
    return (await getCategoriesRequest()).data;
}
