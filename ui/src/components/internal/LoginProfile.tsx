import { useAuth0 } from "@auth0/auth0-react";
import { Avatar, AvatarImage } from "@/components/ui/avatar";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";

export const LoginProfile = () => {
  const { user, isAuthenticated, isLoading, logout } = useAuth0();

  if (isLoading) {
    return <Skeleton className="h-12 w-12 rounded-full" />;
  }

  return (
    isAuthenticated && (
      <Popover>
        <PopoverTrigger asChild>
          <Avatar className="cursor-pointer rounded-full">
            <AvatarImage src={user?.picture} alt="user" />
          </Avatar>
        </PopoverTrigger>
        <PopoverContent className="w-40 p-0">
          <Button
            variant="ghost"
            className="focus:border-none hover:bg-transparent w-full text-lg h-12"
            onClick={() =>
              logout({
                logoutParams: { returnTo: `${window.location.origin}/login` },
              })
            }
          >
            Logout
          </Button>
        </PopoverContent>
      </Popover>
    )
  );
};
