import { format } from "date-fns";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/lib/components/ui/table";

type ReviewsTableProps = {
  isLoading: boolean;
  reviews: AppReview[];
};

export function ReviewsTable({ isLoading, reviews }: ReviewsTableProps) {
  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="w-full">
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Title</TableHead>
              <TableHead>Content</TableHead>
              <TableHead>Author</TableHead>
              <TableHead>Rating</TableHead>
              <TableHead>Submitted At</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {reviews.length > 0 ? (
              reviews.map((review, index) => (
                <TableRow key={index}>
                  <TableCell className="truncate max-w-[300px]">
                    {review.title}
                  </TableCell>
                  <TableCell className="truncate max-w-[300px]">
                    {review.content}
                  </TableCell>
                  <TableCell>{review.author}</TableCell>
                  <TableCell>{review.rating}/5</TableCell>
                  <TableCell>
                    {format(new Date(review.submitted_at), "PPpp")}
                  </TableCell>
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell colSpan={5} className="h-24 text-center">
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
