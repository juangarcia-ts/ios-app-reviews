import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import { HomePage } from "./pages/HomePage";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReviewsPage } from "./pages/ReviewsPage";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/apps" element={<HomePage />} />
          <Route path="/apps/:appId" element={<ReviewsPage />} />
          <Route path="/apps/:appId/reviews" element={<ReviewsPage />} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  );
}

export default App;
