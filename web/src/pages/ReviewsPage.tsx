import { keepPreviousData, useQuery } from "@tanstack/react-query";
import { ReviewsTable } from "../components/ReviewsTable";
import { useEffect, useState } from "react";
import { useParams } from "react-router";
import { PageLayout } from "../components/PageLayout";
import { Text } from "@/lib/components/ui/text";
import { queryClient } from "../query-client";
import { getMonitoredApp } from "../api/monitoredApps";
import { getPaginatedAppReviews } from "../api/appReviews";

export const ReviewsPage = () => {
  const { appId } = useParams();
  const [currentPage, setCurrentPage] = useState(1);

  const { data: app, isLoading: isAppLoading } = useQuery({
    queryKey: ["app"],
    enabled: !!appId,
    queryFn: () => getMonitoredApp(appId!),
  });

  const {
    data: paginatedReviews,
    isPlaceholderData,
    isLoading: isReviewsLoading,
  } = useQuery({
    queryKey: ["reviews", currentPage],
    enabled: !!app?.app_id,
    placeholderData: keepPreviousData,
    queryFn: () => getPaginatedAppReviews(appId!, currentPage, 20),
  });

  // Prefetch the next page due to server-side pagination
  // Ref: https://tanstack.com/query/latest/docs/framework/react/examples/pagination
  useEffect(() => {
    if (!isPlaceholderData && paginatedReviews?.hasMore) {
      queryClient.prefetchQuery({
        queryKey: ["reviews", paginatedReviews.page + 1],
        queryFn: () =>
          getPaginatedAppReviews(appId!, paginatedReviews.page + 1, 20),
      });
    }
  }, [paginatedReviews, isPlaceholderData, appId, queryClient]);

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  if (isAppLoading || isReviewsLoading) {
    return null;
  }

  return (
    <PageLayout
      breadcrumbs={[
        { label: "Home", href: "/" },
        ...(app
          ? [
              {
                label: app.nickname || app.app_name,
                href: `/apps/${appId}/reviews`,
              },
            ]
          : []),
      ]}
    >
      {app ? (
        <>
          <Text.H1 className="mb-8">
            {app.nickname || app.app_name}'s reviews
          </Text.H1>
          <Text.P>
            Reviews from past 48 hours. Reviews are synced every 5 minutes.
          </Text.P>
          <ReviewsTable
            isLoading={isReviewsLoading}
            reviews={paginatedReviews?.data || []}
            currentPage={paginatedReviews?.page || 1}
            totalPageCount={paginatedReviews?.totalPages || 1}
            onPageChange={handlePageChange}
          />
        </>
      ) : (
        <Text.H1>App not found</Text.H1>
      )}
    </PageLayout>
  );
};
