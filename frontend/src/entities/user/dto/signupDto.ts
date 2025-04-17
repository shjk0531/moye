// src/entities/user/dto/signupDto.ts

export interface SignupRequest {
    email: string;
    password: string;
    nickname: string;
}

export interface SignupResponse {
    message: string;
}
