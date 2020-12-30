package controllers

import (
	`appserver/v0/models`
	`context`
	`encoding/json`
	`go.mongodb.org/mongo-driver/bson`
	`net/http`
	`strings`
	`time`
)

// SignInUser Used for Signing In the Users
func SignInUser(response http.ResponseWriter, request *http.Request) {
	var loginRequest models.LoginParams
	var result models.UserDetails
	var errorResponse = models.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&loginRequest)
	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {
		errorResponse.Code = http.StatusBadRequest
		if loginRequest.Email == "" {
			errorResponse.Message = "Email can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if loginRequest.Password == "" {
			errorResponse.Message = "Password can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {

			collection := Client.Database("test").Collection("users")

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var err = collection.FindOne(ctx, bson.M{
				"email":    loginRequest.Email,
				"password": loginRequest.Password,
			}).Decode(&result)

			defer cancel()

			if err != nil {
				returnErrorResponse(response, request, errorResponse)
			} else {
				tokenString, _ := CreateJWT(loginRequest.Email)

				if tokenString == "" {
					returnErrorResponse(response, request, errorResponse)
				}

				var successResponse = models.SuccessResponse{
					Code:    http.StatusOK,
					Message: "You are registered, login again",
					Response: models.SuccessfulLoginResponse{
						AuthToken: tokenString,
						Email:     loginRequest.Email,
					},
				}

				successJSONResponse, jsonError := json.Marshal(successResponse)

				if jsonError != nil {
					returnErrorResponse(response, request, errorResponse)
				}
				response.Header().Set("Content-Type", "application/json")
				response.Write(successJSONResponse)
			}
		}
	}
}

// SignUpUser Used for Signing up the Users
func SignUpUser(response http.ResponseWriter, request *http.Request) {
	var registationRequest models.RegistationParams
	var errorResponse = models.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&registationRequest)
	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {
		errorResponse.Code = http.StatusBadRequest
		if registationRequest.Username == "" {
			errorResponse.Message = "Username can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if registationRequest.Email == "" {
			errorResponse.Message = "Email can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if registationRequest.Password == "" {
			errorResponse.Message = "Password can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {
			tokenString, _ := CreateJWT(registationRequest.Email)

			if tokenString == "" {
				returnErrorResponse(response, request, errorResponse)
			}

			var registrationResponse = models.SuccessfulLoginResponse{
				AuthToken: tokenString,
				Email:     registationRequest.Email,
			}

			collection := Client.Database("test").Collection("users")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, databaseErr := collection.InsertOne(ctx, bson.M{
				"email":    registationRequest.Email,
				"password": registationRequest.Password,
				"username":     registationRequest.Username,
			})
			defer cancel()

			if databaseErr != nil {
				returnErrorResponse(response, request, errorResponse)
			}

			var successResponse = models.SuccessResponse{
				Code:     http.StatusOK,
				Message:  "You are registered, proceed to login",
				Response: registrationResponse,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				returnErrorResponse(response, request, errorResponse)
			}
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(successResponse.Code)
			response.Write(successJSONResponse)
		}
	}
}

// GetUserDetails Used for getting the user details using user token
func GetUserDetails(response http.ResponseWriter, request *http.Request) {
	var result models.UserDetails
	var errorResponse = models.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	bearerToken := request.Header.Get("Authorization")
	var authorizationToken = strings.Split(bearerToken, " ")[1]

	email, _ := VerifyToken(authorizationToken)
	if email == "" {
		returnErrorResponse(response, request, errorResponse)
	} else {
		collection := Client.Database("test").Collection("users")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var err = collection.FindOne(ctx, bson.M{
			"email": email,
		}).Decode(&result)

		defer cancel()

		if err != nil {
			returnErrorResponse(response, request, errorResponse)
		} else {
			var successResponse = models.SuccessResponse{
				Code:     http.StatusOK,
				Message:  "You are logged in successfully",
				Response: result.Username,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				returnErrorResponse(response, request, errorResponse)
			}
			response.Header().Set("Content-Type", "application/json")
			response.Write(successJSONResponse)
		}
	}
}
func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage models.ErrorResponse) {
	httpResponse := &models.ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
	panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}