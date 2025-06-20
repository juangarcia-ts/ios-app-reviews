import {
  Breadcrumb,
  BreadcrumbLink,
  BreadcrumbSeparator,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage,
} from "@/lib/components/ui/breadcrumb";
import { Fragment } from "react/jsx-runtime";

type BreadcrumbItem = {
  label: string;
  href: string;
};

type PageLayoutProps = {
  breadcrumbs?: BreadcrumbItem[];
  children: React.ReactNode;
};

export const PageLayout = ({ children, breadcrumbs }: PageLayoutProps) => {
  return (
    <main className="flex h-full w-full px-4 py-8">
      <div className="container mx-auto w-full max-w-7xl">
        {breadcrumbs && (
          <Breadcrumb className="mb-4">
            <BreadcrumbList>
              {breadcrumbs.map((item, index) => (
                <Fragment key={item.label}>
                  {index > 0 && <BreadcrumbSeparator />}
                  <BreadcrumbItem key={item.label}>
                    <BreadcrumbLink href={item.href}>
                      {item.label}
                    </BreadcrumbLink>
                  </BreadcrumbItem>
                </Fragment>
              ))}
            </BreadcrumbList>
          </Breadcrumb>
        )}

        {children}
      </div>
    </main>
  );
};
