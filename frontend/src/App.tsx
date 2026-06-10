import { FlameIcon } from "lucide-react";

export default function App() {
  return (
        <div className="w-screen h-screen flex items-center justify-center">
            <h1 className="text-3xl font-bold flex items-center gap-4">
                <FlameIcon size={64} />
                <span>Tasks</span>
            </h1>
        </div>
    )
}