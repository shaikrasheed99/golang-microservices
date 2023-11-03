export type RequestBody = {
    action: String,
    signup?: {
      username: String,
      password: String,
      email: String
    },
    login?:{
        username: String,
        password: String
    }
}
  
export type Response = {
    status: String,
    code: String,
    message: String,
    reason?: String,
    data?: any
}

export const convertRequestBodyToFormattedCode = (request: RequestBody): String => {
    let formattedCode: String = "";

    if (request.action === "auth_signup") {
        formattedCode = `
    {
        "action": ${request.action};
        "signup": {
            "username": ${request.signup?.username};
            "password": ${request.signup?.password};
            "email": ${request.signup?.email};
        }
    }`;
    } else if (request.action === "auth_login") {
        formattedCode = `
    {
        "action": ${request.action};
        "login": {
            "username": ${request.login?.username};
            "password": ${request.login?.password};
        }
    }`;
    }

    return formattedCode;
};

export const convertResponseToFormattedCode = (response: Response): String => {
    let formattedCode: String = "";

    if (response.status === "success") {
        formattedCode = `
    {
        "status": ${response.status},
        "code": ${response.code},
        "message": ${response.message},
        "data": ${response.data}
    }`;
    } else if (response.status === "error") {
        formattedCode = `
    {
        "status": ${response.status},
        "code": ${response.code},
        "message": ${response.message},
        "reason": ${response.reason}
    }`;
    }

    return formattedCode;
};