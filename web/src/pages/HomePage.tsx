import { useMutation, useQuery } from "@tanstack/react-query";
import { PageLayout } from "../components/PageLayout";
import { Text } from "@/lib/components/ui/text";
import { MonitoredApp } from "../types";
import { MonitoredAppCard } from "../components/MonitoredAppCard";
import {
  createMonitoredApp,
  deleteMonitoredApp,
  findAllMonitoredApps,
} from "../api/monitoredApps";
import { NewMonitoredAppDialog } from "../components/NewMonitoredAppDialog";
import { queryClient } from "../query-client";

export const HomePage = () => {
  const { data: apps, isLoading } = useQuery({
    queryKey: ["apps"],
    queryFn: findAllMonitoredApps,
  });

  const { mutate: createMonitoredAppMutation } = useMutation({
    mutationFn: createMonitoredApp,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["apps"] });
    },
  });

  const { mutate: deleteMonitoredAppMutation } = useMutation({
    mutationFn: deleteMonitoredApp,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["apps"] });
    },
  });

  return (
    <PageLayout breadcrumbs={[{ label: "Home", href: "/" }]}>
      <div className="flex flex-row justify-between">
        <div className="flex flex-col gap-4">
          <Text.H1>Monitored Apps</Text.H1>
          <Text.P>
            Below is a list of apps currently being monitored. Reviews for these
            apps are automatically fetched from the App Store on a regular
            basis.
          </Text.P>
        </div>
        <NewMonitoredAppDialog onSubmit={createMonitoredAppMutation} />
      </div>

      <div className="mt-8">
        {isLoading ? (
          <div>Loading...</div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {apps?.map((app: MonitoredApp) => (
              <MonitoredAppCard
                key={app.app_id}
                app={app}
                onDelete={deleteMonitoredAppMutation}
              />
            ))}
          </div>
        )}
      </div>
    </PageLayout>
  );
};
