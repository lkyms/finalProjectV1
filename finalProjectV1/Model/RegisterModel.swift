//
//  RegisterModel.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/23.
//

import Foundation

struct jsonValida: Codable {
    var phone: String
}

struct jsonRegis: Codable {
    var nickname: String
    var realname: String
    var username: String
    var password: String
    var phone: String
    var Cap: String
}

struct jsonLog: Codable {
    var Username: String
    var Password: String
}

//struct jsonPost: Codable {
//
//    var Message: String
//    var Status: Int
//    var token: String
////        "Message": "密码正确",
////        "Status": 0,
////        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiMTIzIiwicGhvbmUiOiIzIiwiZXhwIjoxNjM3NzIyNjgzLCJpc3MiOiJhZG1pbi1zZXJ2ZXIifQ.Kq_CM6tNsJuHUwRA4p5_I_of4iQ0p0k4ZpaNvYWbOaY"
////
//}
