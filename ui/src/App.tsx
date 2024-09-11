import { useState } from "react";
import { Button } from "@/components/ui/button";

function App() {
  const [count, setCount] = useState(0);
  return (
    <div className="h-screen min-w-[450px]">
      <div className="h-full xl:mx-[10%] sm:mx-[5%] flex bg-green-200 p-1">
        <Button onClick={() => setCount((prev) => prev + 1)}>
          Count is: {count}
        </Button>
      </div>
    </div>
  );
}

export default App;
