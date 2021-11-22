//
//  LoginViewModel.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/21.
//

import Foundation

protocol LoginViewModel{
    func checkUserNameAndPassword(_ name: String, _ pswd: String) -> Bool
}

class LoginViewModelImpl: ObservableObject, LoginViewModel {
    
    var service: LoginService
    
    init(service: LoginService) {
        self.service = service
    }
    
    func checkUserNameAndPassword(_ name: String, _ pswd: String) -> Bool {
        print(service.checkUserNameAndPassword(name, pswd))
        return service.checkUserNameAndPassword(name, pswd)
    }
    
    
}
