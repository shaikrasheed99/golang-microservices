import { RequestBody, Response, convertRequestBodyToFormattedCode, convertResponseToFormattedCode } from "../utils/helper";
import "./APIContent.css"

interface APIContentProps {
    requestBody: RequestBody | undefined,
    response: Response | undefined
}

const APIContent = ({ requestBody, response }: APIContentProps) => {
    return (
        <div className="api-content">
            <div className="box request-body">
                <p className="title">Request Body</p>
                <hr />
                <pre className="formatted-code">
                    {requestBody === undefined && <p className="empty">Didn't send request yet!</p>}
                    {requestBody !== undefined && <code>{convertRequestBodyToFormattedCode(requestBody)}</code>}
                </pre>
            </div>

            <div className="box response">
                <p className="title">Response</p>
                <hr />
                <pre className="formatted-code">
                    {response === undefined && <p className="empty">No response yet!</p>}
                    {response !== undefined && <code>{convertResponseToFormattedCode(response)}</code>}
                </pre>
            </div>
        </div>
    );
}

export default APIContent;