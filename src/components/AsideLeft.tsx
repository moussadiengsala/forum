
import {  BookmarkIcon, ClipboardDocumentIcon, HomeIcon, MagnifyingGlassIcon, UserGroupIcon, UserIcon, XMarkIcon } from "@heroicons/react/24/outline";
import { useState } from "react";
import { createPortal } from "react-dom";
import { Link, NavLink } from "react-router-dom"
import CreattePost from "./CreattePost";

type Nav = {
    title: string;
    icon: React.ForwardRefExoticComponent<Omit<React.SVGProps<SVGSVGElement>, "ref"> & {
        title?: string | undefined;
        titleId?: string | undefined;
    } & React.RefAttributes<SVGSVGElement>>;
    href: string
}

function AsideLeft() {
    let [isOpenModal, SetIsOpenModal] = useState(false)
    const navigation: Nav[] = [
        {title: 'Home', href: "home", icon: HomeIcon },
        {title: 'Explore', href: "explore", icon: MagnifyingGlassIcon},
        {title: 'List', href: "list", icon: ClipboardDocumentIcon},
        {title: 'Save', href: "save", icon: BookmarkIcon},
        {title: 'Communities', href: "communities", icon: UserGroupIcon},
        {title: 'Profile', href: "#", icon: UserIcon},
    ]

    return (
        <>
            <aside className="bg-yellow-500 h-screen">
                <nav>
                    {navigation.map(nav => (
                        <li key={nav.title}>
                            <NavLink to={nav.href} className={({isActive}) => isActive ? "bg-red-500":"bg-tranparent" }>
                                <nav.icon className="w-8" />
                                <span>{nav.title}</span>
                            </NavLink>
                        </li>
                    ))}
                </nav>
                <button onClick={() => SetIsOpenModal(val => !val)} className="w-full h-10 bg-blue-500">
                    <span>Poster</span>
                </button>
            </aside>

            {isOpenModal && createPortal(
                <div className="absolute w-full h-full left-0 top-0 bg-black/50 flex justify-center items-center">
                    <div className="w-1/2 h-fit bg-red-500 rounded-lg flex justify-center items-center relative">
                        <button onClick={() => SetIsOpenModal(val => !val)}  className="absolute top-4 left-4 w-8">
                            <XMarkIcon />
                        </button>
                        <div>
                            <span>P</span>
                        </div>
                        <CreattePost />
                    </div>
                </div>,
                document.body
            )}
        </>
    )
}

export default AsideLeft