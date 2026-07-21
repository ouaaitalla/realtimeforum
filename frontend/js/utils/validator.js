export function validateRequired(value, fieldName) {

    if (!value || value.trim() === "") {
        return `${fieldName} is required`;
    }

    return null;
}


export function validateEmail(email) {

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;


    if (!emailRegex.test(email)) {
        return "Invalid email format";
    }

    return null;
}


export function validatePassword(password) {

    if (password.length < 6) {
        return "Password must contain at least 6 characters";
    }

    return null;
}


export function validateAge(age) {

    if (age === "") {
        return null;
    }


    const numberAge = Number(age);


    if (isNaN(numberAge)) {
        return "Invalid age";
    }


    if (numberAge < 13) {
        return "Age must be at least 13";
    }


    return null;
}


export function validateRegister(data) {

    const errors = [];


    const nicknameError = validateRequired(
        data.nickname,
        "Nickname"
    );

    if (nicknameError) {
        errors.push(nicknameError);
    }


    const firstNameError = validateRequired(
        data.first_name,
        "First name"
    );

    if (firstNameError) {
        errors.push(firstNameError);
    }


    const lastNameError = validateRequired(
        data.last_name,
        "Last name"
    );

    if (lastNameError) {
        errors.push(lastNameError);
    }


    const emailError = validateEmail(data.email);

    if (emailError) {
        errors.push(emailError);
    }


    const passwordError = validatePassword(data.password);

    if (passwordError) {
        errors.push(passwordError);
    }


    const ageError = validateAge(data.age);

    if (ageError) {
        errors.push(ageError);
    }

    if (!data.gender) {
        errors.push("Please select a gender.");
    }


    return errors;
}

export function validateLogin(data) {

    const errors = [];


    const identifierError = validateRequired(
        data.email,
        "Email or Nickname"
    );

    if (identifierError) {
        errors.push(identifierError);
    } else {

        // If it looks like an email (contains @), validate the format
        if (data.email.includes("@")) {

            const emailFormatError = validateEmail(data.email);

            if (emailFormatError) {
                errors.push(emailFormatError);
            }

        }
        // Otherwise treat it as a nickname — no format validation needed
    }


    const passwordError = validateRequired(
        data.password,
        "Password"
    );

    if (passwordError) {
        errors.push(passwordError);
    }


    return errors;
}

