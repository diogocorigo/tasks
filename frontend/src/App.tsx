import { FlameIcon } from "lucide-react";

export default function App() {
  return (
        <div className="w-screen h-screen flex items-center justify-center">
            <h1 className="text-3xl font-bold flex items-center gap-2">
                <FlameIcon size={32} />
                <span>Tasks</span>
            </h1>
        </div>
    )
}