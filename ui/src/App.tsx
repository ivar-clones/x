import { LoginProfile } from "./components/internal/LoginProfile";
import { useAuth0 } from "@auth0/auth0-react";
import { Icons } from "@/components/ui/icons";
import { Button } from "@/components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { Outlet, useLocation, useNavigate } from "react-router-dom";
import { useMutation, useQuery } from "@tanstack/react-query";
import axios from "axios";
import { useEffect, useState } from "react";

function App() {
  const { isAuthenticated, user } = useAuth0();
  const navigate = useNavigate();
  const location = useLocation();
  const [selectedTab, setSelectedTab] = useState<"for you" | "following">(
    "for you"
  );

  const {
    data: currentUser,
    refetch: fetchUser,
    isFetched: isUserFetched,
  } = useQuery({
    queryKey: ["user"],
    queryFn: () =>
      axios
        .get(`http://localhost:3000/api/v1/users/${user?.email}`)
        .then((res) => {
          if (res.status === 204) {
            return null;
          }
          return res.data;
        }),
    enabled: false,
  });

  const { mutate: createUser } = useMutation({
    mutationFn: (data: {
      name: string | undefined;
      email: string | undefined;
    }) =>
      axios
        .post("http://localhost:3000/api/v1/users", data)
        .then((res) => res.data),
  });

  useEffect(() => {
    if (isAuthenticated) {
      fetchUser();
    }
  }, [isAuthenticated]);

  useEffect(() => {
    if (user && currentUser === null && isUserFetched && isAuthenticated) {
      createUser({
        name: user.name,
        email: user.email,
      });
    }
  }, [user, currentUser, isAuthenticated]);

  if (!isAuthenticated) {
    navigate("/login");
    return;
  }

  return (
    <div className="h-screen min-w-[450px]">
      <div className="h-full flex flex-row">
        <div className="flex flex-col-reverse w-20 md:w-96 items-center py-5 px-2 gap-4">
          {isAuthenticated ? (
            <>
              <LoginProfile />
              <Button
                className="rounded-full w-[80%] h-12 dark:text-white font-bold mb-10"
                size="icon"
                onClick={() => navigate("/post")}
              >
                <Icons.plus className="block md:hidden text-lg " />
                <span className="text-lg hidden md:block">Post</span>
              </Button>
              <Tooltip>
                <TooltipTrigger asChild>
                  <Button
                    variant="ghost"
                    className="rounded-full h-12"
                    onClick={() => navigate("/profile")}
                  >
                    <Icons.user />
                    <span className="text-lg hidden md:block ml-2">
                      Profile
                    </span>
                  </Button>
                </TooltipTrigger>
                <TooltipContent>
                  <p>Profile</p>
                </TooltipContent>
              </Tooltip>
            </>
          ) : undefined}
        </div>
        <div className="flex flex-col overflow-y-auto w-full pr-0 md:pr-10 items-start border-l-2">
          <div className="flex flex-col items-stretch w-full lg:w-[80%] xl:w-[60%] border-r-2 px-2 pb-2">
            <div className="sticky top-0 h-12 w-full flex flex-row gap-5 items-center backdrop-blur mb-2">
              {location.pathname.includes("profile") ? (
                <>
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => navigate(-1)}
                  >
                    <Icons.back />
                  </Button>
                  <div className="flex flex-col gap-2">
                    <h2>{user?.name}</h2>
                  </div>
                </>
              ) : (
                <div className="flex flex-row w-full justify-between">
                  <Button
                    variant="ghost"
                    className={`w-full ${
                      selectedTab === "for you"
                        ? "underline decoration-primary decoration-4 underline-offset-[12px]"
                        : ""
                    }`}
                    size="lg"
                    onClick={() => setSelectedTab("for you")}
                  >
                    <span className="text-lg">For you</span>
                  </Button>
                  <Button
                    variant="ghost"
                    className={`w-full ${
                      selectedTab === "following"
                        ? "underline decoration-primary decoration-4 underline-offset-[12px]"
                        : ""
                    }`}
                    size="lg"
                    onClick={() => setSelectedTab("following")}
                  >
                    <span className="text-lg">Following</span>
                  </Button>
                </div>
              )}
            </div>
            <Outlet />
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
