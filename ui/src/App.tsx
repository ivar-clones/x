import { LoginProfile } from "./components/internal/LoginProfile";
import { useAuth0 } from "@auth0/auth0-react";
import { Icons } from "@/components/ui/icons";
import { Button } from "@/components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";

function App() {
  const { isAuthenticated } = useAuth0();

  return (
    <div className="h-screen min-w-[450px]">
      <div className="h-full flex flex-row">
        <div className="flex flex-col-reverse w-20 md:w-56 items-center py-5 px-2 gap-4">
          {isAuthenticated ? (
            <>
              <LoginProfile />
              <Tooltip>
                <TooltipTrigger asChild>
                  <Button variant="ghost" className="rounded-full">
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
          <div className="flex flex-col items-stretch w-full lg:w-[80%] xl:w-[60%] gap-4 border-r-2 p-2">
            <div className="min-h-44 w-full bg-green-100 text-background rounded-md p-2">
              testing 1
            </div>
            <div className="min-h-44 w-full bg-green-100 text-background rounded-md p-2">
              testing 2
            </div>
            <div className="min-h-44 w-full bg-green-100 text-background rounded-md p-2">
              testing 3
            </div>
            <div className="min-h-44 w-full bg-green-100 text-background rounded-md p-2">
              testing 4
            </div>
            <div className="min-h-44 w-full bg-green-100 text-background rounded-md p-2">
              testing 5
            </div>
            <div className="min-h-44 w-full bg-green-100 text-background rounded-md p-2">
              testing 6
            </div>
            <div className="min-h-44 w-full bg-green-100 text-background rounded-md p-2">
              testing 7
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
