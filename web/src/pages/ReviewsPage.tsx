import { useQuery } from "@tanstack/react-query";
import { ReviewsTable } from "../components/ReviewsTable";
import axios from "axios";
import { useMemo } from "react";
import { useParams } from "react-router";
import { PageLayout } from "../components/PageLayout";
import { Text } from "@/lib/components/ui/text";

const API_URL = import.meta.env.VITE_API_URL;

export const ReviewsPage = () => {
  const { appId } = useParams();

  const { data: app } = useQuery({
    queryKey: ["app"],
    queryFn: () => {
      return axios
        .get(`${API_URL}/api/v1/apps/${appId}`)
        .then((res) => res.data);
    },
  });

  const { data: paginatedReviews, isLoading } = useQuery({
    queryKey: ["reviews"],
    enabled: !!app?.app_id,
    queryFn: () => {
      return axios
        .get(`${API_URL}/api/v1/apps/${appId}/reviews?limit=50`)
        .then((res) => res.data);
    },
  });

  const reviews = useMemo(() => {
    return paginatedReviews?.data || [];
  }, [paginatedReviews]);

  if (!app) {
    return (
      <PageLayout>
        <Text.H1>App not found</Text.H1>
      </PageLayout>
    );
  }

  return (
    <PageLayout
      breadcrumbs={[
        { label: "Home", href: "/" },
        { label: app.nickname, href: `/apps/${appId}/reviews` },
      ]}
    >
      <Text.H1 className="mb-4">{app.nickname}'s reviews ⭐</Text.H1>
      <ReviewsTable isLoading={isLoading} reviews={reviews} />
    </PageLayout>
  );
};
