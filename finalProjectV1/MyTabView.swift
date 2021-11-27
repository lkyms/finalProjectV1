//
//  MyTabView.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/26.
//

import SwiftUI

struct MyTabView: View {
    var body: some View {
        TabView {
            MainPageView()
                .tabItem {
                    Image(systemName: "house")
                        .padding()
                    Text("Home")
                        .font(.title)
                }
            MessageVIew()
                .tabItem {
                    Image(systemName: "message")
                        .padding()
                    Text("Message")
                        .font(.title)
                }
            BoardView()
                .tabItem {
                    Image(systemName: "square.and.pencil")
                        .padding()
                    Text("Board")
                        .font(.title)
                }
            MeView()
                .tabItem {
                    Image(systemName: "person")
                        .padding()
                    Text("Me")
                        .font(.title)
                }
        }
    }
}

struct MyTabView_Previews: PreviewProvider {
    static var previews: some View {
        MyTabView()
    }
}
