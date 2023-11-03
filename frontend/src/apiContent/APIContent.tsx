import "./APIContent.css"

const APIContent = () => {
    let requestBody = `
    {
        "action": "auth_signup",
        "signup": {
            "username": "ironman1",
            "password": "ironman@123",
            "email": "ironman@gmail.com"
        }
    }
    `;

    let response = `
    {
        "status": "success",
        "code": "OK",
        "message": "successfully saved user details",
        "data": null
    }
    `;

    return (
        <div className="api-content">
            <div className="box request-body">
                <p className="title">Request Body</p>
                <hr />
                <pre className="formatted-code">
                    {requestBody === undefined && <p className="empty">Didn't send any request yet!</p>}
                    {requestBody !== undefined && <code>{requestBody}</code>}
                </pre>
            </div>

            <div className="box response">
                <p className="title">Response</p>
                <hr />
                <pre className="formatted-code">
                    {response === undefined && <p className="empty">No response yet!</p>}
                    {response !== undefined && <code>{response}</code>}
                </pre>
            </div>
        </div>
    );
}
 
export default APIContent;