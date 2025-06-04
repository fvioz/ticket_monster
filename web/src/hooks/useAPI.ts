import { useEffect, useState } from "react";

export const useAPI = <T>(
  root: string,
  apiPath: string,
): { data: T | undefined; loading: boolean; error: string | null } => {
  const [value, setValue] = useState<T | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:8080" + apiPath);
        if (!response.ok) {
          throw new Error("Failed to fetch events");
        }
        // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access
        const data = (await response.json())[root] as T;

        setValue(data);
      } catch (err) {
        if (err instanceof Error) {
          setError(err.message);
        } else {
          setError("An unknown error occurred");
        }
      } finally {
        setLoading(false);
      }
    };

    void fetchData();
  }, []);

  return { data: value, loading, error };
};
