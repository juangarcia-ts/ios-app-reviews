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

type MonitoredAppCardProps = {
  app: MonitoredApp;
};

export const MonitoredAppCard = ({ app }: MonitoredAppCardProps) => {
  const navigate = useNavigate();

  const viewReviews = (appId: string) => {
    navigate(`/apps/${appId}/reviews`);
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>{app.nickname}</CardTitle>
      </CardHeader>
      <CardContent>
        <CardDescription>App Store ID: {app.app_id}</CardDescription>
        <CardDescription>
          Last synced: {format(new Date(app.last_synced_at), "PPpp")}
        </CardDescription>
      </CardContent>
      <CardFooter>
        <CardAction>
          <Button className="pointer" onClick={() => viewReviews(app.app_id)}>
            View reviews
          </Button>
        </CardAction>
      </CardFooter>
    </Card>
  );
};
