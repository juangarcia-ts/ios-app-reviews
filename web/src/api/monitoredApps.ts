import axios from "axios";
import { AppInfo } from "../types";

const API_URL = import.meta.env.VITE_API_URL;

export const findAllMonitoredApps = async () => {
  return axios.get(`${API_URL}/api/v1/apps`).then((res) => res.data);
};

export const getMonitoredApp = async (appId: string) => {
  return axios.get(`${API_URL}/api/v1/apps/${appId}`).then((res) => res.data);
};

export const createMonitoredApp = async (data: {
  appId: string;
  appName: string;
  logoUrl: string;
  nickname: string | undefined;
}) => {
  return axios
    .post(`${API_URL}/api/v1/apps`, {
      app_id: data.appId,
      app_name: data.appName,
      logo_url: data.logoUrl,
      nickname: data.nickname,
    })
    .then((res) => res.data);
};

export const deleteMonitoredApp = async (appId: string) => {
  return axios
    .delete(`${API_URL}/api/v1/apps/${appId}`)
    .then((res) => res.data);
};

export const getAppInfoFromStore = async (appId: string): Promise<AppInfo> => {
  return axios
    .get(`${API_URL}/api/v1/apps/${appId}/lookup`)
    .then((res) => res.data);
};

export const syncReviews = async (appId: string) => {
  return axios
    .post(`${API_URL}/api/v1/apps/${appId}/sync`)
    .then((res) => res.data);
};
