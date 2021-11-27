//
//  LoginService.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/21.
//

import Foundation

protocol LoginService {
    func checkUserNameAndPassword(_ name: String, _ pswd: String) -> Bool
}

struct LoginServiceImpl: LoginService {
    
    var model = CheckLogin()
    var result = jsonPostSuccess(Message: "", Status: -1, token: "")
    var flag1 = -1
    
    func checkUserNameAndPassword(_ name: String, _ pswd: String) -> Bool {
//        model.postDataLog()
        model.checkUserNameAndPassword(name, pswd) {  result in
            
            
            return (result == 1)
//
//
//
//            return false
        }
        
        if flag1 == 1 {
            return true
        } else {
            return false
        }
    }
    
    
}
