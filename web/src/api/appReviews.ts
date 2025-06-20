import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;

export const getPaginatedAppReviews = async (
  appId: string,
  page: number,
  limit: number
) => {
  return axios
    .get(`${API_URL}/api/v1/apps/${appId}/reviews?page=${page}&limit=${limit}`)
    .then((res) => res.data);
};
