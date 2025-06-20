export interface AppReview {
  id: string;
  app_id: string;
  title: string;
  content: string;
  author: string;
  rating: number;
  submitted_at: string;
  created_at: string;
  updated_at: string;
}

export interface MonitoredApp {
  app_id: string;
  app_name: string;
  logo_url: string;
  nickname: string | null;
  last_synced_at: string;
  created_at: string;
  updated_at: string;
}

export interface AppInfo {
  app_name: string;
  logo_url: string;
}
