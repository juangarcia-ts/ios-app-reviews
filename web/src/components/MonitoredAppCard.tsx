import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/lib/components/ui/card";
import { MonitoredApp } from "../types";
import { format } from "date-fns";
import { Button } from "@/lib/components/ui/button";
import { useNavigate } from "react-router";
import { RefreshCcwIcon, Trash2Icon } from "lucide-react";
import { useCallback } from "react";
import { syncReviews } from "../api/monitoredApps";
import { queryClient } from "../query-client";
import { useMutation } from "@tanstack/react-query";

type MonitoredAppCardProps = {
  app: MonitoredApp;
  onDelete: (appId: string) => void;
};

export const MonitoredAppCard = ({ app, onDelete }: MonitoredAppCardProps) => {
  const navigate = useNavigate();

  const { mutate: syncReviewsMutation } = useMutation({
    mutationFn: syncReviews,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["apps"] });
    },
  });

  const viewReviews = useCallback(
    (appId: string) => {
      navigate(`/apps/${appId}/reviews`);
    },
    [navigate]
  );

  return (
    <Card>
      <CardHeader>
        <div className="flex flex-row items-center justify-between">
          <CardTitle className="text-lg">
            {app.nickname || app.app_name}
          </CardTitle>
          <Button
            variant="outline"
            size="icon"
            onClick={() => onDelete(app.app_id)}
          >
            <Trash2Icon className="w-4 h-4" />
          </Button>
        </div>
        <img src={app.logo_url} alt={app.app_name} className="w-10 h-10" />
      </CardHeader>
      <CardContent>
        <CardDescription>App Store ID: {app.app_id}</CardDescription>
        <CardDescription>Status: Healthy</CardDescription>
        <CardDescription>
          Last synced:{" "}
          {app.last_synced_at
            ? format(new Date(app.last_synced_at), "PPpp")
            : "Never"}
        </CardDescription>
      </CardContent>
      <CardFooter>
        <CardAction className="flex flex-row gap-2">
          <Button className="pointer" onClick={() => viewReviews(app.app_id)}>
            View reviews
          </Button>
          <Button
            variant="outline"
            onClick={() => syncReviewsMutation(app.app_id)}
          >
            <RefreshCcwIcon className="w-4 h-4" /> Sync now
          </Button>
        </CardAction>
      </CardFooter>
    </Card>
  );
};
