import { useQuery } from "@tanstack/react-query";
import { ReviewsTable } from "../components/ReviewsTable";
import axios from "axios";
import { useMemo } from "react";

const API_URL = import.meta.env.VITE_API_URL;

export const HomePage = () => {
  const { data, isLoading } = useQuery({
    queryKey: ["reviews"],
    queryFn: () => {
      return axios
        .get(`${API_URL}/api/v1/reviews?appId=447188370&limit=50`)
        .then((res) => res.data);
    },
  });

  const reviews = useMemo(() => {
    return data?.data || [];
  }, [data]);

  return (
    <div className="flex h-full w-full items-center justify-center">
      <div className="container mx-auto w-full max-w-7xl">
        <ReviewsTable isLoading={isLoading} reviews={reviews} />
      </div>
    </div>
  );
};
