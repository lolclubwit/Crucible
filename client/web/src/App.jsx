import { useState } from "react";
import SplitPane from "split-pane-react";
import "split-pane-react/esm/themes/default.css";
import Editor from "./components/editor";

export default function App() {
	const [sizes, setSizes] = useState([100, "30%", "auto"]);
	const [sizes2, setSizes2] = useState([100, "30%", "auto"]);

	return (
		<div className="h-[100vh]">
			<SplitPane split="vertical" sizes={sizes} onChange={setSizes} resizerSize={10}>
				<div className="h-full flex justify-center items-center bg-slate-400">
					Question Pane
				</div>
				<div className="h-full">
					<SplitPane split="" sizes={sizes2} onChange={setSizes2} resizerSize={10}>
						<div className="h-full bg-slate-500">
							<Editor />
						</div>
						<div className="h-full flex justify-center items-center bg-slate-300">
							Result Pane
						</div>
					</SplitPane>
				</div>
			</SplitPane>
		</div>
	);
}
