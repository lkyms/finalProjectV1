//
//  LoginService.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/21.
//

import Foundation

protocol LoginService {
    func checkUserNameAndPassword(_ name: String, _ pswd: String) //-> Bool
}

struct LoginServiceImpl: LoginService {
    
    var model = CheckLogin()
    
    func checkUserNameAndPassword(_ name: String, _ pswd: String) {//-> Bool {
        model.postDataLog()
        //model.checkUserNameAndPassword(name, pswd)
    }
    
    
}
