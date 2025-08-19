async function trylogin(){
    const inputLogin = document.getElementById("InputLogin");
    const inputPassword = document.getElementById("InputPassword");

    const user_input = {
        login: inputLogin,
        password: inputPassword
    }


    //API
    try{
       const response = await fetch('http://localhost:1323/', {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json',
            },
            body: JSON.stringify(user_input)
        })
        if (!response.ok){
            throw new Error(`HTTP error! status: ${response.status}`);
        }


        const responseData = await response.json();
        if (!responseData.success) {
            throw new Error(responseData.message || 'Server response indicates failure');
        }

        return responseData;

        } catch (error) {
            console.error('Failed to send user data:', error);
            throw error;
        }
}

async function handleLoginButtonClick() {
  try {
    const result = await trylogin();
    console.log("Login successful:", result);
  } catch (error) {
    console.error("Login failed:", error.message);
  }
}