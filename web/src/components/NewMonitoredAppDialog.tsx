import { Button } from "@/lib/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/lib/components/ui/dialog";
import { Input } from "@/lib/components/ui/input";
import { Label } from "@/lib/components/ui/label";
import { DialogOverlay } from "@radix-ui/react-dialog";
import { useCallback, useEffect, useState } from "react";
import { getAppInfoFromStore } from "../api/monitoredApps";
import { useQuery } from "@tanstack/react-query";

type NewMonitoredAppDialogProps = {
  onSubmit: (app: {
    appId: string;
    appName: string;
    logoUrl: string;
    nickname: string | undefined;
  }) => void;
};

export const NewMonitoredAppDialog = ({
  onSubmit,
}: NewMonitoredAppDialogProps) => {
  const [appId, setAppId] = useState("");
  const [appName, setAppName] = useState("");
  const [nickname, setNickname] = useState("");
  const [logoUrl, setLogoUrl] = useState("");
  const [isOpen, setIsOpen] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const { data: appInfo, refetch: refetchAppInfo } = useQuery({
    queryKey: ["lookup", appId],
    queryFn: () => getAppInfoFromStore(appId),
    enabled: false,
    retry: false,
  });

  const handleSubmit = useCallback(async () => {
    setIsSubmitting(true);

    try {
      onSubmit({ appId, appName, logoUrl, nickname });
      setAppId("");
      setNickname("");
      setIsOpen(false);
    } catch (error) {
      console.error("Failed to create monitored app:", error);
    } finally {
      setIsSubmitting(false);
    }
  }, [appId, appName, nickname, onSubmit]);

  useEffect(() => {
    if (appId) {
      refetchAppInfo();
    }
  }, [appId, refetchAppInfo]);

  useEffect(() => {
    if (appInfo) {
      setAppName(appInfo.app_name);
      setLogoUrl(appInfo.logo_url);
    } else {
      setAppName("");
      setLogoUrl("");
    }
  }, [appInfo]);

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button variant="outline">Monitor New App</Button>
      </DialogTrigger>
      <DialogOverlay />
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Monitor New App</DialogTitle>
          <DialogDescription>
            Enter the App Store ID of the app you want to monitor.
          </DialogDescription>
        </DialogHeader>

        <div className="grid gap-4">
          <div className="grid gap-3">
            <Label htmlFor="app-id">App Store ID</Label>
            <Input
              required
              id="app-id"
              name="app-id"
              placeholder="1234567890"
              value={appId}
              onChange={(e) => {
                setAppId(e.target.value);
              }}
            />
          </div>
          <div className="grid gap-3">
            <Label htmlFor="app-name">App name</Label>
            <Input
              required
              id="app-name"
              name="app-name"
              placeholder="App Name (from App Store)"
              value={appName}
              disabled
            />
          </div>
          <div className="grid gap-3">
            <Label htmlFor="nickname">Nickname</Label>
            <Input
              required
              id="nickname"
              name="nickname"
              placeholder="Nickname"
              value={nickname}
              onChange={(e) => setNickname(e.target.value)}
            />
          </div>
        </div>

        <DialogFooter>
          <DialogClose asChild>
            <Button type="button" variant="outline">
              Cancel
            </Button>
          </DialogClose>
          <Button
            type="button"
            onClick={handleSubmit}
            disabled={!appInfo || !appName || isSubmitting}
          >
            {isSubmitting ? "Saving..." : "Save changes"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};
