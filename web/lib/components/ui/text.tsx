import { cn } from "@/lib/utils";

const H1 = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <h1
      className={cn(
        "scroll-m-20 text-3xl font-extrabold tracking-tight text-balance",
        className
      )}
    >
      {children}
    </h1>
  );
};

const H2 = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <h2
      className={cn(
        "scroll-m-20 text-2xl font-extrabold tracking-tight text-balance",
        className
      )}
    >
      {children}
    </h2>
  );
};

const H3 = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <h3
      className={cn(
        "scroll-m-20 text-xl font-extrabold tracking-tight text-balance",
        className
      )}
    >
      {children}
    </h3>
  );
};

const H4 = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <h4
      className={cn(
        "scroll-m-20 text-lg font-extrabold tracking-tight text-balance",
        className
      )}
    >
      {children}
    </h4>
  );
};

const P = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <p className={cn("text-sm text-muted-foreground", className)}>{children}</p>
  );
};

const Span = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <span className={cn("text-sm text-muted-foreground", className)}>
      {children}
    </span>
  );
};

export const Text = {
  H1,
  H2,
  H3,
  H4,
  P,
  Span,
};
