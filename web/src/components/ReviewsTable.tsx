import { format } from "date-fns";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/lib/components/ui/table";
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { AppReview } from "../types";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/lib/components/ui/tooltip";
import {
  Pagination,
  PaginationNext,
  PaginationEllipsis,
  PaginationLink,
  PaginationItem,
  PaginationPrevious,
  PaginationContent,
} from "@/lib/components/ui/pagination";
import { Text } from "@/lib/components/ui/text";

type ReviewsTableProps = {
  isLoading: boolean;
  reviews: AppReview[];
  currentPage: number;
  totalPageCount: number;
  onPageChange: (page: number) => void;
};

export function ReviewsTable({
  isLoading,
  reviews,
  currentPage,
  totalPageCount,
  onPageChange,
}: ReviewsTableProps) {
  const columns: ColumnDef<AppReview>[] = [
    {
      header: "Title",
      accessorKey: "title",
      cell: ({ row }) => {
        return (
          <Tooltip>
            <TooltipTrigger className="truncate max-w-[200px]">
              {row.original.title}
            </TooltipTrigger>
            <TooltipContent className="max-w-[500px]">
              {row.original.title}
            </TooltipContent>
          </Tooltip>
        );
      },
    },
    {
      header: "Content",
      accessorKey: "content",
      cell: ({ row }) => {
        return (
          <Tooltip>
            <TooltipTrigger className="truncate max-w-[500px]">
              {row.original.content}
            </TooltipTrigger>
            <TooltipContent className="max-w-[500px]">
              {row.original.content}
            </TooltipContent>
          </Tooltip>
        );
      },
    },
    {
      header: "Author",
      accessorKey: "author",
    },
    {
      header: "Rating",
      accessorKey: "rating",
      cell: ({ row }) => {
        return (
          <div className="flex items-center gap-2">
            {Array.from({ length: row.original.rating }).map((_, index) => (
              <span key={index} className="text-yellow-500">
                ⭐
              </span>
            ))}
          </div>
        );
      },
    },
    {
      header: "Submitted At",
      accessorKey: "submitted_at",
      cell: ({ row }) => {
        return format(new Date(row.original.submitted_at), "PPpp");
      },
    },
  ];

  if (isLoading) {
    return <Text.P>Loading...</Text.P>;
  }

  const table = useReactTable({
    data: reviews,
    columns: columns,
    getCoreRowModel: getCoreRowModel(),
  });

  // Always show 5 page numbers, centered on current page when possible
  const getPageNumbers = () => {
    const maxVisiblePages = 5;
    const pages = [];

    let start = Math.max(
      1,
      Math.min(
        currentPage - Math.floor(maxVisiblePages / 2),
        totalPageCount - maxVisiblePages + 1
      )
    );
    let end = Math.min(totalPageCount, start + maxVisiblePages - 1);

    for (let i = start; i <= end; i++) {
      pages.push(i);
    }

    // If there are less than maxVisiblePages, pad at the start
    while (pages.length < maxVisiblePages && pages[0] > 1) {
      pages.unshift(pages[0] - 1);
    }

    return pages;
  };

  const pageNumbers = getPageNumbers();

  return (
    <div>
      <div className="rounded-md border mb-4">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext()
                          )}
                    </TableHead>
                  );
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <Pagination>
        <PaginationContent>
          <PaginationItem>
            <PaginationPrevious
              onClick={() => onPageChange(currentPage - 1)}
              className={
                currentPage <= 1
                  ? "pointer-events-none opacity-50"
                  : "cursor-pointer"
              }
            />
          </PaginationItem>

          {totalPageCount > 5 && currentPage - 2 > 1 && (
            <>
              <PaginationItem>
                <PaginationLink
                  onClick={() => onPageChange(1)}
                  className="cursor-pointer"
                >
                  1
                </PaginationLink>
              </PaginationItem>
              <PaginationEllipsis />
            </>
          )}

          {pageNumbers.map((pageNum) => (
            <PaginationItem key={pageNum}>
              <PaginationLink
                onClick={() => onPageChange(pageNum)}
                isActive={currentPage === pageNum}
                className="cursor-pointer"
              >
                {pageNum}
              </PaginationLink>
            </PaginationItem>
          ))}

          {totalPageCount > 5 && currentPage + 2 < totalPageCount && (
            <>
              <PaginationEllipsis />
              <PaginationItem>
                <PaginationLink
                  onClick={() => onPageChange(totalPageCount)}
                  className="cursor-pointer"
                >
                  {totalPageCount}
                </PaginationLink>
              </PaginationItem>
            </>
          )}

          <PaginationItem>
            <PaginationNext
              onClick={() => onPageChange(currentPage + 1)}
              className={
                currentPage >= totalPageCount
                  ? "pointer-events-none opacity-50"
                  : "cursor-pointer"
              }
            />
          </PaginationItem>
        </PaginationContent>
      </Pagination>
    </div>
  );
}
