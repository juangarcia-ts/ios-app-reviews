import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;

export const findAllMonitoredApps = async () => {
  return axios.get(`${API_URL}/api/v1/apps`).then((res) => res.data);
};

export const getMonitoredApp = async (appId: string) => {
  return axios.get(`${API_URL}/api/v1/apps/${appId}`).then((res) => res.data);
};
