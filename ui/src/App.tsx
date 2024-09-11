import { LoginButton } from "./components/internal/LoginButton";
import { LoginProfile } from "./components/internal/LoginProfile";
import { useAuth0 } from "@auth0/auth0-react";

function App() {
  const { isAuthenticated } = useAuth0();

  return (
    <div className="h-screen min-w-[450px]">
      <div className="h-full xl:mx-[10%] sm:mx-[5%] flex p-2 flex-row gap-2">
        <div className="flex flex-col-reverse min-w-12 w-12 xl:w-72 lg:w-64 md:w-52 sm:w-44 items-center py-5">
          {!isAuthenticated && <LoginButton />}
          {isAuthenticated && <LoginProfile />}
        </div>
        <div className="flex flex-col overflow-y-auto w-full gap-10">
          <div className="min-h-44 w-full xl:w-[70%] lg:w-96 bg-green-100 text-background rounded-md p-2">
            testing 1
          </div>
          <div className="min-h-44 w-full xl:w-[70%] lg:w-96 bg-green-100 text-background rounded-md p-2">
            testing 2
          </div>
          <div className="min-h-44 w-full xl:w-[70%] lg:w-96 bg-green-100 text-background rounded-md p-2">
            testing 3
          </div>
          <div className="min-h-44 w-full xl:w-[70%] lg:w-96 bg-green-100 text-background rounded-md p-2">
            testing 4
          </div>
          <div className="min-h-44 w-full xl:w-[70%] lg:w-96 bg-green-100 text-background rounded-md p-2">
            testing 5
          </div>
          <div className="min-h-44 w-full xl:w-[70%] lg:w-96 bg-green-100 text-background rounded-md p-2">
            testing 6
          </div>
          <div className="min-h-44 w-full xl:w-[70%] lg:w-96 bg-green-100 text-background rounded-md p-2">
            testing 7
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
