import { useAuth0 } from "@auth0/auth0-react";

export const RedirectLogin = () => {
  const { isAuthenticated, loginWithRedirect } = useAuth0();

  if (!isAuthenticated) {
    loginWithRedirect();
    return;
  }

  return <div></div>;
};
