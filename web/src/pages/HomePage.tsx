import { useQuery } from "@tanstack/react-query";
import { ReviewsTable } from "../components/ReviewsTable";
import axios from "axios";
import { useMemo } from "react";
import { PageLayout } from "../components/PageLayout";
import { Text } from "@/lib/components/ui/text";
import { MonitoredApp } from "../types";
import { MonitoredAppCard } from "../components/MonitoredAppCard";

const API_URL = import.meta.env.VITE_API_URL;

export const HomePage = () => {
  const { data: apps, isLoading } = useQuery({
    queryKey: ["reviews"],
    queryFn: () => {
      return axios.get(`${API_URL}/api/v1/apps`).then((res) => res.data);
    },
  });

  return (
    <PageLayout breadcrumbs={[{ label: "Home", href: "/" }]}>
      <div className="flex flex-col gap-4">
        <Text.H1>Monitored Apps 📲</Text.H1>
        <Text.P>
          Below is a list of apps currently being monitored. Reviews for these
          apps are automatically fetched from the App Store on a regular basis.
        </Text.P>
      </div>

      <div className="mt-8">
        {isLoading ? (
          <div>Loading...</div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {apps?.map((app: MonitoredApp) => (
              <MonitoredAppCard key={app.app_id} app={app} />
            ))}
          </div>
        )}
      </div>
    </PageLayout>
  );
};
